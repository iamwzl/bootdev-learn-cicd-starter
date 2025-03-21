#  Boot.dev -- 21. Learn CI/CD
![CI workflow status](https://github.com/iamwzl/bootdev-learn-cicd-starter/actions/workflows/ci.yml/badge.svg)    ![CD workflow status](https://github.com/iamwzl/bootdev-learn-cicd-starter/actions/workflows/cd.yml/badge.svg)

![A logo representing git actions](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/pcgPbf9.png)

My work as part of the boot.dev Learn CI/CD course.

[Check out boot.dev (with a referal from me)](https://wzl.to/boot.dev)  ***-or-*** [Check it out, without a referal](https://wzl.to/boot.dev_noref)

- - -
# learn-cicd-starter (Notely)

This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

## Local Development

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

Run the server:

```bash
go build -o notely && ./notely
```

*This starts the server in non-database mode.* It will serve a simple webpage at `http://localhost:8080`.

You do *not* need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!


Doing A Lesson's version of the Boot.dev's Notely App.
