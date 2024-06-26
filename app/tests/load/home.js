import http from "k6/http";
import { check } from "k6";

// Test configuration
export const options = {
  thresholds: {
    // Assert that 99% of requests finish within 3000ms.
    http_req_duration: ["p(99) < 3000"],
  },
  // Ramp the number of virtual users up and down
  stages: [
    { duration: "10s", target: 10 },
    // { duration: "1m", target: 40 },
    // { duration: "10s", target: 0 },
  ],
};

// Simulated user behavior
export default function () {
  let res = http.get("https://example.com");
  // Validate response status
  check(res, { "status was 200": (r) => r.status == 200 });
}
