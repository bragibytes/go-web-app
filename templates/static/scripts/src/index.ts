import authentication from "./controllers/authentication"
import users from "./controllers/users";
import posts from "./controllers/posts";
import voting from "./controllers/voting";
import comments from "./controllers/comments";

declare var $: any;

$(document).ready(function() {
    $(".dropdown-trigger").dropdown();
});

console.log("handling....")
authentication()
users()
posts()
comments()
voting()



