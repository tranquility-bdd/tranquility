[![Codacy Badge](https://api.codacy.com/project/badge/Grade/5a8c1957079146108d37c090167f1f58)](https://app.codacy.com/app/filfreire/tranquility-bdd?utm_source=github.com&utm_medium=referral&utm_content=tranquility-bdd/tranquility-bdd&utm_campaign=Badge_Grade_Dashboard)
[![CircleCI](https://circleci.com/gh/tranquility-bdd/tranquility-bdd.svg?style=svg)](https://circleci.com/gh/tranquility-bdd/tranquility-bdd)

# tranquility
a lightweight dependency/bloat-free alternative to serenity-rest-assured and serenity-core

## The mission
- **maintainable** - be more maintainable than postman - with little clutter - but equally as fast to develop full-length check suites as postman
- **featherweight** - little dependencies (contrasting to serenity 170+ dependencies, or newmans 140+ dependencies)
- **flexible** - do not lock the user to a "BDD" approach where user is forced to use Gherkin in order to use `tranquility`

## Feature roadmap (2 June 2019)
- *(MVP)* Users can code their actions and manipulate environment (which influences actions)
- *(MVP)* Create example repository that uses MVP version of tranquility with `gucumber`
- Users can use PreAction and Test template/abstractions as part of their tests, if they opt for a unit approach
- Users can define "Lego piece" of the 3 main components - PreAction, Action and Test
- (Optional/Doubt/Nice-to-have) Be able to export Scenarios built with Lego pieces to JSONs that obey postman schema
