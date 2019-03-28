[![CircleCI](https://circleci.com/gh/tranquility-bdd/tranquility-bdd.svg?style=svg)](https://circleci.com/gh/tranquility-bdd/tranquility-bdd)

# tranquility-bdd
a lightweight dependency/bloat-free alternative to serenity-rest-assured and serenity-core

## How to build

Using maven (version 3, Java 8):
```
mvn clean install
```

## Desired features (work in progress)

- ? extend Cucumber runtime (similar to CucumberWithSerenity)
- mimic serenity’s RequestSpecification class (maybe? maybe not needed - users could simply use restassured)
- mimic “setting environment variables” like Serenity
- support cucumber reports
- return html reports (with clear request, response and curl requests) similar to the ones of serenity (simple ones, plain html, no javascript clutter, maybe tacit css for making it less bland than normal html)
- option to log requests in curl format (in case of rest assured flavour)  during test execution
- suport gradle and maven projects
