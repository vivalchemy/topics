let ls = document.getElementById("ls");
let ss = document.getElementById("ss");
let c = document.getElementById("c");

// accessing the storage locations
// it is always available. stays even when the tab or browser is closed
// it doesn't have any expiry date an need to be manually removed or the browser to be closed
// not sent with the requests
console.log("local storage");
localStorage.setItem("name", "Vivian");
localStorage.setItem("removable", "item");
console.log(localStorage.getItem("name"));
console.log("Before removal from ls", localStorage.getItem("removable"));
localStorage.removeItem("removable");
console.log("Before removal from ls", localStorage.getItem("removable")); // this returns null

// session storage is only avaiable for the duration that tab is open in the current browser. Once the tab of browser is closed the session storage data is lost
// it doesn't have any expiry date an need to be manually removed or the browser to be closed
// not sent with the requests
console.log("session storage");
sessionStorage.setItem("name", "Vivian");
sessionStorage.setItem("removable", "item");
console.log(sessionStorage.getItem("name"));
console.log("Before removal from ss", sessionStorage.getItem("removable"));
sessionStorage.removeItem("removable");
console.log("Before removal from ss", sessionStorage.getItem("removable")); // this returns null

// they are sent with the request
// they are automatically deleted since we have to set the expiry date
// only valid before the expiration date
console.log("cookies");
document.cookie = "name=vivial";
let data = "vanilla";
document.cookie = "yummy_cookie=choco";
document.cookie = `tasty_cookie=${data}`;
document.cookie = `name=Vivian; expires=${new Date(2024, 7, 15).toUTCString()}`; // for some weird reason the month starts from 0
// this will overwrite the previous cookie
document.cookie = "virat=benstrokes";

// printing on the page
ls.innerHTML = "local storage: " + JSON.stringify(localStorage, null, 2);
ss.innerHTML = "session storage: " + JSON.stringify(sessionStorage, null, 2);
c.innerHTML = "cookie: " + JSON.stringify(document.cookie, null, 2);

// overwriting the value in ls ans ss is just essentially resetting the value
sessionStorage.setItem("name", "Vivian");
sessionStorage.setItem("name", "Vivian Ludrick");

//console.log(new Date());
