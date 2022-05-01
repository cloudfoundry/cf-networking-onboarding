---
layout: single
title: Incoming HTTP Requests Part 5 - Access Logs
permalink: /http-routes/incoming-http-requests-pt-5
sidebar:
  title: "HTTP Routes"
  nav: sidebar-http-routes
---

## Assumptions
- You have a CF deployed
- You have one
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE
- You have completed the previous stories in this track

## Recorded values from previous stories
```
APP_A_ROUTE=<value>
APP_A_GUID=<value>
DIEGO_CELL_IP=<value>
CONTAINER_APP_PORT=<value>
DIEGO_CELL_APP_PORT=<value>
CONTAINER_ENVOY_PORT=<value>
DIEGO_CELL_ENVOY_PORT=<value>
OVERLAY_IP=<value>
```

## What
In the previous stories you followed the path of a request from a client to an
app deployed on Cloud Foundry.

For every successful* response that Gorouter returns to a client it logs an
access log. Here successful means "that there was a response from the app with
any status code". 

These access logs can be very helpful for debugging. One common situation is
that a customer sees via metrics that they are getting lots of 502s. But what
apps are returning 502s? Let's look at the access logs to find out!

## How

📝 **Look at access logs**
1. Ssh onto the Router VM and become root
1. Tail the access log
   ```bash
   tail /var/vcap/sys/log/gorouter/access.log
   ```

1. In another terminal curl APP_A_ROUTE.

   You should see something that looks like this:
   ```
   APP_A_ROUTE - [2021-02-18T21:22:32.355501523Z] "GET / HTTP/1.1" 200 0 62 "-" "curl/7.54.0" "35.191.2.80:51628" "10.0.1.11:61002" x_forwarded_for:"142.105.202.35, 35.227.211.74, 35.191.2.80" x_forwarded_proto:"http" vcap_request_id:"6cb8b8de-bc85-4479-5f90-1c3a52d88d84" response_time:0.011962 gorouter_time:0.000467 app_id:"cabd9e08-384e-4689-b868-1ba3d5d838bf" app_index:"0" x_cf_routererror:"-" x_b3_traceid:"9d30688835227cf3" x_b3_spanid:"9d30688835227cf3" x_b3_parentspanid:"-" b3:"9d30688835227cf3-9d30688835227cf3"
   ```

* ❓Can you find the status code in the access log?
* ❓Can you find the x-cf-routererror in the access log?

📚 **Read about the X-CF-RouterError**
1. Read about the [X-CF-RouterError here](https://docs.cloudfoundry.org/adminguide/troubleshooting-router-error-responses.html#gorouter-specific-response-headers) and learn how it can be used for debugging. 

🤔 **Look at the app logs for APP_A**.
1. Use `cf logs` to look at the app logs for `APP_A`.
 * ❓Can you find a log line that looks like the access log line?
 * ❓What additional information does the app log contain?

## Expected Result
You have found the access log from your curl in the log file on the router VM
and in the app logs.

## Resources
* [X-CF-RouterError
  docs](https://docs.cloudfoundry.org/adminguide/troubleshooting-router-error-responses.html#gorouter-specific-response-headers)
