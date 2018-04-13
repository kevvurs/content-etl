# NewsRoom Project
I am building a web platform for news content similar to news.google.com.
The most significant feature of my webapp will be the ability for users to
post comments externally on the content items and rate the content across
a variety of metrics. Users should be able to engage with news media. The
current polarization of our media signals the need for a stronger feedback
loop between media publications and their readers.

## content-etl
Go microservice for persisting news related content periodically. I am calling
newsapi.org to gather US news headlines and details. The application will
extract, transform, and load the content with some minor revisions.

The application is built to run on Google App Engine (Standard Environment).
For persistence, it stores the data with Goolge Datastore, using an md5 hash
for content ID's. The app has a `cron.yaml` and will eventually be set
to read and store the news data from newsapi.org every 10 minutes or so and
delete aged articles from my datastore.

## Development
Anyone can clone or fork this project to persist data from newsapi.org.
However, there are some things to be aware of first. The runtime is App Engine
specific. That is, running the application _as is_ would require Google Cloud SDK,
particularly the App Engine dev server and the Datastore emulator.

I do not share my `app.yaml` here because it is not relevant for contributing
or cloning. To use this program you also need to sign up for account (free
available) from newsapi.org.
