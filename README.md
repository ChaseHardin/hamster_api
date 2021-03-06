# Hamster API

This project is an end point to harness the Google maps API with Golang. 

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them

-   Install Golang
-   Open `.bash_profile` by typing: `nano .bash_profile`
-   Set go path:

```
#setting go path
export GOROOT="/usr/local/go"
export GOPATH="/Users/userName/go"
export PATH=$GOPATH/bin:$PATH
```

## Git Flow

This project is composed of two branches:

- Master - production
- Dev - test/stage

Developers are to develop features in local branches and make pull requests to merge with the `Dev` branch when ready. 

When code has been accepted it will be merged into the `Master` branch.  

## Deploy API
We are using Heroku to host the hamster Api. These were the steps taken to create/deploy project to Heroku.
API can be found here: https://sheltered-chamber-13682.herokuapp.com/greet

## Step 1
Install Heroku CLI. Install can be found [here](https://devcenter.heroku.com/articles/getting-started-with-go#set-up)

## Step 2
Navigate to the `/app` directory

Install Godep:
```go get github.com/kr/godep```

Save Dependencies:
```godep save```

Add/Commit to git
'''git add app && git commit -m "Initial build/web subtree commit"'''

## Step 3
If this is a new application, create the Heroku app.
```heroku create```

Push code to Heroku

```git subtree push --prefix app heroku master```

## Additional Details
[Getting Started With Go On Heroku](https://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html)