# logger
Simple library used for logging

# Configuration

The [Viper](https://github.com/spf13/viper) library is used for configuration management. 

The configuration is stored inside memory structures, with hardcoded default values. 

It can be adjusted during development using the configuration file, and at deploy time values can be overriden using environment variables. 

The priority of assigned values for setting from lowest to highest priority is the following:
- default values (lowest)
- `.env.file`
- enviromnet variable (highest)

## Configuration options
Changing the default vaues *should be avoided*, but easily can be done directly manipulating the code. 

File `.env.yaml` is *preferd approach* to adjust configuration *during development*. The file is bundled togehter with binary and is part of deployable package.

To override the value at *deploy time* use environment variables. Adjust the name of the variable to the full path to the setting in uppercase by replacing all dots with underscore. 

The follwoing options can be configured:
```
log:
  format: string
  level: debug
``` 
Format can be either one of existing `string` and `json`, or new format can be implemented and used in configuration.
Level can be one of the follwing "panic", "fatal", "error", "warn", "info", "debug" and "trace".

## Example

Default setting assigned in memory structure for log.format value is `string`:

```
    // Default configurtion
    func (l *Config) Default() {
        l.Format = StringFormatter.Name
        l.Level = Error.Name
    }
```

The configuration in `.env.yaml` file for this setting can change it to another:

```
log:
  format: json
  level: debug
```

To change the level to `warn` of the application at deploy time, use the env variable:

```
LOG_LEVEL=warn
``` 

# Usage Exаmple

Formatters for structured logging can be implemented and added easily. 
On global level the configuration defines the format and the minimal logging level for the whole application. 
Configutration is read and applied on global level when the InitLogger functionl is called. If not called, the default values are used.

```
    // -- Initialise configuration
    loggerConfig := new(logger.Config)
    err := loggerConfig.InitConfig()
    if err != nil {
        fmt.Printf("Cannot load the configuration: %s \n", err)
        os.Exit(1)
    }

    // -- Initialise logging system
    err = logger.InitLogger(loggerConfig)
    if err != nil {
        fmt.Printf("Cannot init the logging system: %s \n", err)
        os.Exit(1)
    }
```

Then use the logger for pringing the messages: 

```
    // Create named logger
    mainLogger := logger.NewNamedLogger("main logger") 

    mainLogger.Panicf("%+v", appConfig)
    mainLogger.Panic("Panic entry")
```

