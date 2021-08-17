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

See [openweathermap.yml.example](https://github.com/easyas314159/prometheus-openweathermap/blob/main/openweathermap.yml.example) for an example and detailed documentation on available configuration options.
