//const promiseOne = new Promise(function (resolve, reject) {
//  // Do async calls like network req, db, etc
//  setTimeout(function () {
//    console.log("promiseOne completed");
//  }, 1000);
//  // the promise is not resolved so the promiseOne will not be called
//});

// calling promise one => will output "promiseOne completed" after 1 second
//promiseOne.then(function () {
//  console.log("promiseOne resolved");
//});

//const promiseTwo = new Promise(function (resolve, reject) {
//  // Do async calls like network req, db, etc
//  setTimeout(function () {
//    console.log("promiseTwo completed");
//    when resolve function is called the function in the .then method will be called for it
//    resolve();
//  }, 1000);
//});

//promiseTwo.then(function () {
//  console.log("promiseTwo resolved");
//});

// .then method called be directly called on a promise without assigning it to a new variable
//new Promise(function (resolve, reject) {
//  setTimeout(function () {
//    console.log("promiseThree completed");
//    resolve();
//  }, 1000);
//}).then(function () {
//  console.log("promiseThree resolved");
//});

// object passed in the resolve function can be received as object in the .then callback function parameter
//const promiseFour = new Promise(function (resolve, reject) {
//  setTimeout(function () {
//    console.log("promiseFour completed");
//    resolve({ username: "Vivian", age: 19 });
//  }, 1000);
//});
//
//promiseFour.then(function ({ username, age }) {
//  console.log(username, age);
//  console.log("promiseFour resolved");
//});

// resolve is used to indicate promise success and reject for promise failure
//const promiseFive = new Promise(function (resolve, reject) {
//  setTimeout(function () {
//    let error = true;
//    if (!error) {
//      resolve({ username: "Vivian", age: 19 });
//    }
//    reject("ERROR: Something went wrong");
//  }, 1000);
//});
//
//// the .then method consume the return values of the previous .then method
//promiseFive
//  .then(function (user) {
//    console.log("user ", user);
//    return user.username;
//  })
//  .then(function (username) {
//    console.log("username ", username);
//  })
//  .catch((err) => {
//    console.log(err);
//  })
//  .finally(() => {
//    console.log("finally it is over");
//  });

//const promiseSix = new Promise(function (resolve, reject) {
//  setTimeout(function () {
//    let error = true;
//    if (!error) {
//      resolve({ username: "Async-await", age: 19 });
//    }
//    reject("ERROR: Something went wrong");
//  }, 1000);
//});

// we can use async await instead of the typical .then and .catch methods too
// async await cannot directly handle error so we need to use try catch
//async function display() {
//  try {
//    const user = await promiseSix;
//    console.log(user);
//  } catch (err) {
//    console.log("Graceful ", err);
//  }
//}
//display();

//---------------------------------------------------------------------
// Same code in async await and .then .catch
//---------------------------------------------------------------------
async function getUser(url) {
  try {
    const response = await fetch(url);
    // NOTE: response.json method will take time to be converted so always use await
    const user = await response.json();
    console.log("Got the users via await", user);
  } catch (err) {
    console.log("Graceful try catch", err);
  }
}
getUser("https://jsonplaceholder.typicode.com/users/1");

function display(url) {
  fetch(url)
    // response.json will convert from the response type to the json type
    .then((response) => response.json())
    // the user will then get the response in json format
    .then((user) => console.log("Got the users via .then ", user))
    //catch will then catch if there will be any errors
    .catch((err) => console.log("Graceful .then .catch method ", err));
}
display("https://jsonplaceholder.typicode.com/users/1");
