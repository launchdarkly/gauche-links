function saveOptions() {
  const devHost = $('#dev-host').val();
  const devMode = $('#dev-mode').is(':checked');
  const host = $('#host').val();
  chrome.storage.sync.set({
    devHost: devHost,
    devMode: devMode,
    host: host,
  });
}

function restoreOptions() {
  chrome.storage.sync.get({
    devHost: DefaultDevHost,
    devMode: false,
    host: DefaultHost,
  }, function (items) {
    $('#dev-host').val(items.devHost || DefaultDevHost);
    $('#host').val(items.host || DefaultHost);
    $('#dev-mode').prop('checked', items.devMode);
  });
}

$('input').on('change input', saveOptions);
$(document).ready(() => {
  $('#dev-host').prop('placeholder', DefaultDevHost);
  $('#host').prop('placeholder', DefaultHost);
  restoreOptions();
});
