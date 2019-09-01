// Save this so we can handle blocking requests for webRequest
let SavedHostUrl;

const prefix = Keyword.replace(/\/$/, '');

function getHostUrl(cb) {
  chrome.storage.sync.get({
    devHost: DefaultDevHost,
    devMode: false,
    host: DefaultHost
  }, function (items) {
    let finalHost;
    if (items.devMode) {
      finalHost = items.devHost || DefaultDevHost;
    } else {
      finalHost = items.host || DefaultHost;
    }
    SavedHostUrl = new URL(finalHost);
    if (cb) {
      cb(SavedHostUrl);
    }
  });
}

// Trigger an asynchronous update from local storage
setTimeout(getHostUrl);

chrome.omnibox.onInputStarted.addListener(function () {
  getHostUrl(function(hostUrl) {
    $.ajax(hostUrl.toString() + "_search"); // Wake up the host
  });
  chrome.omnibox.setDefaultSuggestion({ description: "Start typing..." });
});

function getSuggestions(text, cb) {
  getHostUrl(function(hostUrl) {
    $.ajax(hostUrl.toString() + "_search?q=" + text, {dataType: "json"}).done(cb);
  });
}

let currentRequestId = 0;
let lastReceivedRequestId = 0;

// from https://stackoverflow.com/questions/822452/strip-html-from-text-javascript
function stripTags(text) {
  const div = document.createElement("div");
  div.innerHTML = text;
  return div.textContent || div.innerText || "";
}

// from https://stackoverflow.com/questions/9847580/how-to-detect-safari-chrome-ie-firefox-and-opera-browser
const isFirefox = typeof window["InstallTrigger"] !== 'undefined';

// firefox doesn't support <match>, <url> and <dim> tags in descriptions
function formatDescription(description) {
  if (isFirefox) {
    return stripTags(description);
  }
  return description;
}

// This event is fired each time the user updates the text in the omnibox,
// as long as the extension's keyword mode is still active.
chrome.omnibox.onInputChanged.addListener(
  _.throttle((text, suggest) => {
    currentRequestId++;
    const requestId = currentRequestId;
    getSuggestions(text, function (response, textStatus, request) {
      // Just ignore this if it isn't the latest (otherwise these might happen out of order)
      if (requestId < lastReceivedRequestId) {
        return
      }
      lastReceivedRequestId = requestId;
      text = text.trim(); // Ignore lead and trailing whitespace
      const current = `<url>${prefix}/${text}</url>`;
      const results = [];
      if (request.status === 401 || request.status === 302) {
        chrome.omnibox.setDefaultSuggestion({ description: formatDescription("<match>Cannot show suggestions until you log in.  Hit Enter.</match>") });
      } else {
        if (response[1]) {
          let exactMatch;
          for (let i = 0; i < response[1].length; i++) {
            const path = response[1][i];
            const description = response[2][i];
            const pathWithMatch = path.replace(new RegExp(`^(${text})`), "<match>$1</match>");
            const descriptionWithMatch = description.replace(new RegExp(`(${text})`, "i"), "<match>$1</match>");
            const isRegexp = path[0] === "/" && path[path.length-1] === "/";
            let result = {
              content: path,
              description: formatDescription(`<url>${prefix}/${pathWithMatch}</url> · ${descriptionWithMatch}`),
            };
            if (isRegexp) {
              const pattern = path.slice(1, -1);
              const patternWithMatch = pattern.replace(new RegExp(`^(${text})`), "<match>$1</match>");
              result = {
                content: `${'\u200B'.repeat(i)}${extendMatch(text, pattern)}`, // Use zero-width spaces to make these unique
                description: formatDescription(`<dim><url>${prefix}/${patternWithMatch}</url></dim> · ${descriptionWithMatch}`),
              };
            }
            if (path === text && !isRegexp) {
              exactMatch = i;
              results.unshift(result);
            } else {
              results.push(result);
            }
          }
          results.sort(function (a, b) {
            return a.content.length - b.content.length;
          });
          if (exactMatch !== undefined) {
            const description = response[2][exactMatch];
            chrome.omnibox.setDefaultSuggestion({ description: formatDescription(`<match>${current}</match> · ${description}`) });
          } else {
            chrome.omnibox.setDefaultSuggestion({ description: formatDescription(`Add new link <match><url>${current}</url></match>`) });
          }
        } else {
          chrome.omnibox.setDefaultSuggestion({ description: formatDescription(`Add new link <match><url>${current}</url></match>`) });
        }
      }
      suggest(results);
    });
  }, 200, {trailing: true}));

// This event is fired with the user accepts the input in the omnibox.
chrome.omnibox.onInputEntered.addListener(
  function (text, disposition) {
    getHostUrl(function(hostUrl) {
      const strippedText = text.replace(/\u200b/g, "");

      // if this is a regexp partial match, don't redirectNavigation
      if (strippedText != text) {
        // It would be nice if we could do something more useful here but redirecting isn't very helpful
        return
      }

      const url = hostUrl.toString() + strippedText;
      switch (disposition) {
        case "currentTab":
          chrome.tabs.query({ currentWindow: true, active: true }, function (tabs) {
            chrome.tabs.update(tabs[0].id, {url: url});
          });
          break;
        case "newForegroundTab":
          chrome.tabs.create({ url: url, selected: true });
          break;
        case "newBackgroundTab":
          chrome.tabs.create({ url: url, selected: false });
          break;
      }
    })
  }
);

// Extend up to the constant literal prefix of the regular expression
function extendMatch(text, pattern) {
  if (text != pattern.slice(0, text.length)) {
    return text;
  }
  const more = pattern.slice(text.length).match(/[^\\\[\]^${}()?+]*/)[0];
  return text + more;
}

function redirectUrl(url, hostUrl) {
  const destUrl = url;
  destUrl.host = hostUrl.host;
  if (hostUrl.port) {
    destUrl.port = hostUrl.port;
  }
  destUrl.protocol = hostUrl.protocol;
  return destUrl
}

function redirectRequest(event) {
  let url;
  try {
    url = new URL(event.url);
  } catch (e) {
    return
  }

  if (SavedHostUrl) {
    const targetUrl = redirectUrl(url, SavedHostUrl);
    return {redirectUrl: targetUrl.toString()};
  }
}

chrome.storage.onChanged.addListener(function() {
  setTimeout(getHostUrl);
});

const requestUrlsFilter = {urls: ['*://' + prefix + '/*'], types: ["main_frame"]};
chrome.webRequest.onBeforeRequest.addListener(redirectRequest, requestUrlsFilter, ["blocking"]);
