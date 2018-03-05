# LaunchDarkly GaucheLinks

## What is it?

GaucheLinks is a Link Shortening service.  It includes a Chrome Extensions to provide search autocomplete when a prefix is entered in the Chrome Omnibox

## How does it work?

GauchLinks redirects `<host>/<link path>` to a `destination` value loaded from a table.  Links are stored in AirTable containing tuples of `path`, `destination` and `description`.  For example,

`gauchelinks.launchdarkly.com/wiki` -> `https://launchdarkly.atlassian.net/wiki`

## Installing the Chrome Extension

Follow the instructions below to install the latest chrome extension on your browser:

1. Download the latest chrome extension from the [releases](https://github.com/launchdarkly/gauche-links/releases) page.
2. Enter `chrome://extensions` in your browser address bar.
3. Open your downloads folder and drag the extension into the window.

Simply opening the link will not work because of Chrome will not install an extension from outside the Chrome Store.   

## Deployment

The LaunchDarkly GaucheLinks App runs on Google Cloud Platform.  It runs at https://gauche-links.appspot.com/ but we also have a redirection from https://gauche.launchdarkly.com.

To deploy, run:

```
make deploy
```

## Development

To do development you must set up `gcloud` on your host.

```
brew install glcoud
gcloud components install app-engine-go

```

To run the app, you run teh google cloud platform development server:

```make run```
 
Gauche links uses bootstrap 4 for styling.
