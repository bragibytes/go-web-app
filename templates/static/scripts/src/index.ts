import authentication from "./controllers/authentication"
import users from "./controllers/users";
import posts from "./controllers/posts";
import voting from "./controllers/voting";
import comments from "./controllers/comments";
import M from "materialize-css";

document.addEventListener('DOMContentLoaded', function() {
    var elems = document.querySelectorAll('.tabs');
    var instances = M.Tabs.init(elems);
  });


console.log("Javascript is working. You look nice today!")
authentication()
users()
posts()
comments()
voting()



