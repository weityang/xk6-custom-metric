# xk6-custom-metric

This is a [k6](https://go.k6.io/k6) extension using the [xk6](https://github.com/k6io/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
| ---------------------------------------------------------------------------------------------------------------------------- |

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [gvm](https://github.com/moovweb/gvm)
- [Git](https://git-scm.com/)

Then, install [xk6](https://github.com/k6io/xk6) and build your custom k6 binary with the Kafka extension:

1. Install `xk6`:

```shell
$ go install go.k6.io/xk6/cmd/xk6@latest
```

2. Build the binary:

```shell
$ xk6 build --with github.com/weityang/xk6-custom-metric@latest
```

# example

```javascript
import cm from "k6/x/customMetric";
export default function () {
  cm.add(5);
}
```
