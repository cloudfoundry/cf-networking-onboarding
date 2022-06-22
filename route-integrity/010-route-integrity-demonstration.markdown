---
layout: single
title: Route Integrity Demonstration
permalink: /route-integrity/route-integrity-demonstration
sidebar:
  title: "Route Integrity"
  nav: sidebar-route-integrity
---
## What

In the last story you saw the dangers of stale routes. This story will explore
how Route Integrity solves this issue. Route Integrity is enabled by enforcing
TLS (or mTLS) between the GoRouter and the apps it routes to.  By default only
TLS (not mTLS) is enforced. This is because mTLS will break the cf ssh command.

(Side Note: Confusingly, this feature is called "Route Integrity" by engineers,
but is called "(m)TLS to apps" or "encrypting traffic from the GoRouter to
backends" in some documentation. I will continue calling it Route Integrity,
because it is more succinct and sounds cooler.)

Let's set up the same experiment as the previous story and see that even when
you imitate route staleness, that you are no longer able to see the data from
the incorrect app.

## How
🤔 **Send traffic to a stale route**

1. Deploy your CF with Route Integrity enabled with TLS. Use the manifest you
   saved in the last story at `/tmp/env-with-route-integrity.yml`.

1. Push an app, appA with an HTTP route.

1. Push a different app, appB, that will respond differently than appA. Ensure
   there is an HTTP Route mapped to appB. Make sure an instance of appA and
   appB are on the same Diego Cell. You can achieve this by scaling appB, but
   make sure there is only 1 instance of appA.

1. Curl appA. See that you get the expected result.

1. Change the iptables rules so that the Diego Cell port that maps to appA now,
   incorrectly, redirects traffic to appB. This will imitate route staleness.

1. In another terminal, tail the gorouter.stdout.log.

1. Curl the route for appA.

### Expected Result
When you curl appA, you should get a 503 status code.
```bash
$ curl appA.<YOUR DOMAIN>
503 Service Unavailable
```

In the GoRouter bosh logs you should see a line about pruning the bad route.

```json
{"log_level":3,"timestamp":1554761639.745829,"message":"prune-failed-endpoint","source":"vcap.gorouter.registry","data":{"route-endpoint":{"ApplicationId":"bf1cdd21-ed5b-429b-93ef-f74a4aadc55f","Addr":"10.0.1.13:61010","Tags":{"component":"route-emitter"},"RouteServiceUrl":""}}}

{"log_level":3,"timestamp":1554761639.7473993,"message":"backend-endpoint-failed","source":"vcap.gorouter","data":{"route-endpoint":{"ApplicationId":"bf1cdd21-ed5b-429b-93ef-f74a4aadc55f","Addr":"10.0.1.13:61010","Tags":{"component":"route-emitter"},"RouteServiceUrl":""},"error":"x509: certificate is valid for 0c2ca10a-0c19-4288-68e8-8d1f, not e5fdb45c-1b41-4602-45ac-b49e","attempt":1,"vcap_request_id":"3673ccd9-44ee-4d81-4a45-0ea75a96121c"}}
```

Route Integrity doesn't actually prevent route staleness, but it *does* prevent
leaking data from appB when you are trying to access appA.

## Look at the code
GoRouter analyses different types of errors from the backend. In the case of misrouting, we are running into this [HostnameMismatch error](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/fails/basic_classifiers.go#L62-L69).

Depending on the type of error, GoRouter may choose to retry or not. Look at these [failure classifiers](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/fails/classifier_group.go#L13-L22).

## ❓ Questions
1. Are HostnameMismatch errors retried?
1. If you had multiple instances of appA what do think would've happened? (Extra Credit: Try it out. Were you correct?)
1. Route Integrity was introduced to fix misrouting due to route staleness. What other benefits does Route Integrity provide?
1. Keep curling the HTTP route for appA several times over 1 minute. Sometimes you get a 503 and sometimes you get a 404. Why?

## Resources
* [CF Docs - TLS to Apps](https://docs.cloudfoundry.org/concepts/http-routing.html#tls-to-back-end)
