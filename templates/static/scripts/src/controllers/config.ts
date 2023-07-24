import {user} from "../interfaces";

class go_data {
    private go_user:user|null = null

    constructor() {
    }
    get_user(){
        return this.go_user

    }
    set_user(u:user){
        this.go_user = u
    }
}

export const store = new go_data()

export const element_exists = (n:string):boolean => {
    const ele = document.getElementById(n)
    return ele == null ? false:true
}





