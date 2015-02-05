var express = require('express'),
    timers  = require('timers'),
    logger  = require('morgan');

var app = express();

var sleepMiddleware = function(req, res, next){
  console.log('Sloth is sleeping, please wait');
  timers.setTimeout(function(){
    next();
  }, 3000);
}

var teaMiddleware = function(req, res, next){
  console.log('*drinks hibiscus tea*');
  timers.setTimeout(function(){
    next();
  }, 500);
}

var serveSloth = function(req, res){
  res.send('<img src="http://andyhaskell.github.io/Slothful-Soda/images/sloth.jpg" width="240px" height="300px" />');
}



app.use(logger('common'));

app.use('/sloths', sleepMiddleware, function(req, res){
  res.send('<img src="http://andyhaskell.github.io/Slothful-Soda/images/sloth.jpg" width="240px" height="300px" />');
});

/*There was nothing new to demonstrate in the /sloths2 route in the Express
 *version since /sloths2 was added in the Go version to demonstrate how to do
 *middleware chaining that isn't hard-coded, which is how /sloths worked in
 *the Express version.
 */

app.use('/sloths3', sleepMiddleware, teaMiddleware, serveSloth);



app.use('/sleepTeaSlothChain', sleepMiddleware, teaMiddleware, serveSloth)
app.use('/teaSleepSlothChain', teaMiddleware, sleepMiddleware, serveSloth)
app.use('/teaTwiceChain', teaMiddleware, teaMiddleware, serveSloth)



app.use('/', function(req, res){
  res.send('Hello world!');
});

app.listen(1123);
