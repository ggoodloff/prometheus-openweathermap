# prometheus-openweathermap

A daemon process for aggregating current openweathermap.org conditions into prometheus

## Usage

```
$> openweathermap
```

## Configuration

The openweathermap metric aggregator looks for configuration file in the following locations:

1. Current execution environment
2. `.`
3. `${HOME}`
4. `/etc`

### `openweathermap.yml` Example

See [openweathermap.yml.example](openweathermap.yml.example) for an example and detailed documentation on available configuration options.

## Alerting

See the [examples/](examples/) folder in this repository for a docker based example setup that uses prometheus and alertmanager to notify an IoT device via HTTP.
