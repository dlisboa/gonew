# testing

The `testing/` directory is for acceptance tests, mostly for ensuring the
critical user paths in the app work deploy-to-deploy. Avoid adding very specific
tests here, those are better suited for unit tests. These tests preferentially
should also be able function as smoke tests.

By default there's a `testing/browser/` directory, for [Playwright](https://playwright.dev/)
acceptance-style tests with real browsers, and `testing/load/` for load testing
with [K6](https://k6.io/open-source).

Another possibility is having API acceptance tests in a `testing/api/` directory
using [Hurl](https://hurl.dev/), though that is not created automatically.
