# Points of Interest Launcher
Google App Engine initialization for my points of interest back-end server.

###Goal of this project
This project injects the following dependencies into the back-end server,
in order to initialize it in a Google App Engine context:
- `AppEngineLog` implementing `github.com/BoSunesen/pointsofinterest/webapi/logging.Logger`
- `AppEngineContextFactory` implementing `github.com/BoSunesen/pointsofinterest/webapi/factories.ContextFactory`
- `AppEngineClientFactory` implementing `github.com/BoSunesen/pointsofinterest/webapi/factories.ClientFactory`
- `AppengineWorkerFactory` implementing `github.com/BoSunesen/pointsofinterest/webapi/factories.WorkerFactory`

###Links
The main server code can be found here:
- https://github.com/BoSunesen/pointsofinterest

The hosted application can be found here:
- Back-end: https://points-of-interest-1308.appspot.com/poi
- Front-end: https://points-of-interest-map.appspot.com

###My experience with Google App Engine
I have never worked with Google App Engine before. In the end I found that the
ease of deployment was appreciable, though the environment is somewhat restrictive.
Separating "google.golang.org/appengine" dependencies from the main project seems to be essential,
to avoid becoming incompatible with other hosting options.
