import authentication from "./controllers/authentication"
import users from "./controllers/users";
import {element_exists, store} from "./controllers/config";

if(element_exists("hidden-package")){
    const x = document.getElementById("hidden-package") as HTMLInputElement
    const s = x.value
    const u = JSON.parse(s)
    store.set_user(u)
}

console.log("handling....")
authentication()
users()


