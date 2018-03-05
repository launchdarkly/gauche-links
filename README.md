# LaunchDarkly GaucheLinks

## What is it?

GaucheLinks is a Link Shortening service similar to GoLinks (aka Google Short Links) that was formerly available as part of Google's Gsuite.  It includes Chrome and Firefox Extensions to provide search autocomplete when a prefix is entered in the Chrome Omnibox.  The extensions allow you to redirect short links in your browser, such as `http://ld/wiki` to `https://launchdarkly.atlassian.net/wiki`.  They also allow you to search for text in the link and description when you've entered `ld` as your keyword.  The keyword choice is configurable.

## How does it work?

GauchLinks redirects `<host>/<link path>` to a `target` value loaded from a table.  Links are stored in an AirTable base with a table called `Links` containing tuples of `path`, `target`, `description` and `author`.  For example,

`ld.launchdarkly.com/wiki` -> `https://launchdarkly.atlassian.net/wiki`

### Regexp Patterns
GaucheLinks also supports regexp subsitution for paths, which is indicated by surround the path with `/`.  For example, one of our rules is `/pulls/(\w+)/` -> `github.com/launchdarkly/$1/pulls`.  This allows us to see ll the pull requests for a particular repo.

### Prefix Patterns
GaucheLinks will also perform prefix matches, so the rule `gh/` -> `github.com/launchdarkly/` would allow `https://gh/guache-links` to redirect to `https://github.com/launchdarkly/gauche-links`.

### Fragments and Query strings

If the *target* does not specify a fragment or query string, those values will be taken from the user-provided path.  Currently, they are not combined, although it is a possible future enhancement.

## Configuration

Here's an example configuration:

```
GAUCHE_PREFIX: "ld"
GAUCHE_AIRTABLE_APIKEY: "<some api key>"
GAUCHE_AIRTABLE_APPPATH: "/appy02Y7Rb632dTvj"
GAUCHE_AIRTABLE_BASEID: "appy02Y7Rb632dTvj"
GAUCHE_AIRTABLE_BASEPATH: "/tbltdnv863UO1sCi3/viwgDayBfgbAaXZaW"
GAUCHE_AIRTABLE_ROOTURL: "https://airtable.com/"
```

Create an Airtable Base with a table called `Links` and the following schema:

```
Path - Single Line Text
Target - Single Line Text
Description - Single Line Text
Author - Email
```

Then you can construct the configuration above.

## Building the Chrome or Firefox Extension

GaucheLinks takes can be run via the cli to build extensions for Chrome and Firefox.

```guache-links -extenions -prefix <prefix> -extensions-path <where to put it> -host <the host name> -dev-host <the development host name> -version <extension version>```

The sample demo appengine deploy shows how the artifacts can be made available as static links.

## The Database

The currently supported database backend is AirTable.  Support for Google Sheets as a backend is on the roadmap.  In fact, the original name of the project is a portmanteau of "Go Links" and "Google Sheets", but AirTable ended up being more convenient.  Unfortunately, AirTable API Keys are per user and shared across all Bases and their shared forms are public, which may make Google Sheets a more desirable backend.

## Special link names

The GaucheLinks website expects several special links for navigation.  You should add these to your link database.  They are listed below:

```
_add - Add a new database.  In our deployment, this links to a secret form in AirTable that enters a new value in the links table.
_edit - Edit all the links.  This links to a private view of AirTable.
_chrome_extension - The current Chrome extension .crx file
_firefox_extension - The current Firefox extension .fpi file.
```

## Deployment

The GaucheLinks is currently organized to be easily deployable on the Google Cloud Platform.  There is code in the `demo/appengine` to show how to run it.  In our case at LaunchDarkly, we have a repo dedicated to deploying GaucheLinks via CircleCI, that starts with the name of our selected prefix.  Here's a reop that shows how you might deploy GuacheLinks with prefix `eg`:

```https://github.com/launchdarkly/eg-gauche-links```

## Development

To do development you must set up `gcloud` on your host.

```
brew install glcoud
gcloud components install app-engine-go

```

To run the app using the google cloud platform development server:

```cd demo/appengine && make run```
 
Gauche links uses (https://getbootstrap.com/)[Bootstrap 4] for styling and (https://fontawesome.com/)[FontAwesome] for icons.
