[![Circle CI](https://circleci.com/gh/PotomacInnovation/compass.svg?style=svg&circle-token=3f8ea1d3a744d61244aa65d2940734a8b700ae2c)](https://circleci.com/gh/PotomacInnovation/compass)

# Compass
## Make sure you have Meteor installed.
## Clone repo, `npm install` then...
### To run locally, you need to have the Upside Host Bastion pem key at `~/pem-keys/upside.pem`.
### run `sh start.sh`
#### How-to instructions on that part:
```
Steps to Connect to Postgres Database

1. Setup LastPass (this is super fun…)
2. Ask me to share bastion host PEM key with you in LastPass

3. Create file in your `~/` file called `pem-keys`
4. Create file in `~/pem-keys/` called `upside.pem` and copy paste the PEM KEY from LastPass into it.  [can use command line by calling ` cat > upside.pem <paste key>`]

5. On command line, cd into `~/pem-keys/` folder
6. Run `chmod 400 upside.pem`

7. To turn on the tunnel, run `ssh -i ~/pem-keys/upside.pem -f -N -n -o PasswordAuthentication=no -L 9000:develop-upside-engine.cs6uzey25k6n.us-west-2.rds.amazonaws.com:5432 ec2-user@52.26.134.141`

8. In pgAdmin, set host to `localhost` && port to `9000`
9. Set Username to Upside_Engine
10. Enter password
11. Connect

// Core DB
ssh -i ~/pem-keys/upside.pem -f -N -n -o PasswordAuthentication=no -L 9001:develop-core-engine.cs6uzey25k6n.us-west-2.rds.amazonaws.com:5432 ec2-user@52.26.134.141
```
The start shell script has the tunneling part built in, but requires the pem key to function.


### ELB Addresses
| environment | url |
| ----------- | --- |
| dev | https://dev.uscompass.co   |
| int | https://int.uscompass.co   |
| stg | https://stage.uscompass.co |
| prd | https://www.uscompass.co   |

### This application is built with [Meteor](http://docs.meteor.com/#/full/) and [React](https://facebook.github.io/react/docs/getting-started.html) and written in ES2015.
It is linted by [ESLint](http://eslint.org/) using the [AirBnB](https://www.npmjs.com/package/eslint-config-airbnb)
rules with [ESLint React Plugin](https://github.com/yannickcr/eslint-plugin-react).

This application began as a Meteor 1.2 application and moved into a 1.3 application. Meteor 1.3
introduced real NPM packages as well as ES2015 modules. Meteor used to throw around global variables
to share code between files, but now you can import and export. You will see both of these ideas in
the application. You will see patterns like `const SomeComponent = window.SomeComponent` as well as
`export default SomeComponent`.

Meteor now can use native NPM packages as well as [Atmosphere](https://atmospherejs.com/) packages.
You will see both of these in this application as well. You can find the atmostphere packages
in `/.meteor/packages`. If you see something like
`import { FlowRouter } from 'meteor/kadira:flow-router-ssr'` that means it is an atmosphere package
and the author is kadira and the package is flow-router-ssr.

Important packages include [FlowRouter](https://atmospherejs.com/kadira/flow-router-ssr) which
follows familiar routing patterns.

[numtel:pg](https://atmospherejs.com/numtel/pg) which provides reactive subscriptions from Postgres
to the client-side applications through publishing queries. You'll find these in `/server/publish.js`.

[react-meteor-data](https://github.com/meteor/react-packages/blob/47504c7e87649ca5c6be7acb4c3ec57b02e66713/docs/meteor-data.md)
is used mostly as a plugin inside React Components. This was an older way of doing things and mixins
will be disappearing completely from React in v0.15.0 which may already be released. This provided
a reactive data interface in the `getMeteorData(){}` function of a React Class. This means any data
referred to inside the render method which came from that function will cause a re-render when that
data changes. This is the heart of the Meteor Data Reactivity.

The new way to deal with this is still with [react-meteor-data *new link](http://guide.meteor.com/react.html#data),
but with React containers. The container component subscribes to the data, then passes is to a child
component as props. This will cause all of the happy re-rendering that the mixin provided, while
separating the data subscription from the UI component. This way you can have a stateless ui
component that only renders html and a container which only knows how to get data. This makes testing
easier.

You'll find this particular pattern used in `/client/modules/home/bottomMiddleBox/PersonCaseHistoryComponent.jsx`.
This uses a different library from react-meteor-data because I found it easier to use. The idea is the
exact same though.

[Accounts](http://guide.meteor.com/accounts.html#accounts-password) manages user resources. This
includes logins, sessions, hashes of passwords and profiles. This alone might be worth keeping mongo
for.

[HTTP](http://docs.meteor.com/#/full/http) is your fairly average AJAX object, but it is important
to know that it is a core Meteor module.

[Email](http://docs.meteor.com/#/full/email) is the module currently used to send emails through
Mandril using Evan's credentials. This requires an EMAIL_URL environment variable to be exposed
upon startup. Production should be exposing this and it is part of the `start.sh` script locally.

### Meteor things to know
Meteor does a lot for you. Locally, when developing it will refresh your browser and rebundle your
files upon saving. You don't have to worry about any of this stuff. No gulp, no webpack, no worries.

That said, the server-side of things is where it gets weird. Meteor's server wants to look like a
synchronous thread, even if it runs asynchronously. If you are using traditional callbacks, you may
need to make use of `Meteor.wrapAsync()` which will turn your asynchronous callback code into code
that you can call and assign synchronously. This will make the callback finish behind the scenes and
the program will keep running like a normal node application, but it will look like traditional
synchronous code. This takes some getting used to and you may see errors related to futures/fibers
if you have callbacks in your server code that aren't handled properly.

When de-meteorized, this is just a node application. 
