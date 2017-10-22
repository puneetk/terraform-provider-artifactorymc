---
layout: "null"
page_title: "Provider: Artifactory MC"
sidebar_current: "docs-artifactorymc-index"
description: |-
  The artifactory mc provider provides no-op constructs that can be useful helpers in tricky cases.
---

# Artifactory MC Provider

The `artifactory mc` provider is a rather-unusual provider that has constructs that
intentionally do nothing. This may sound strange, and indeed these constructs
do not need to be used in most cases, but they can be useful in various
situations to help orchestrate tricky behavior or work around limitations.

The documentation of each feature of this provider, accessible via the
navigation, gives examples of situations where these constructs may prove
useful.

Usage of the `artifactory mc` provider can make a Terraform configuration harder to
understand. While it can be useful in certain cases, it should be applied with
care and other solutions preferred when available.
