"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.element_exists = exports.store = void 0;
class go_data {
    constructor() {
        this.go_user = null;
    }
    get_user() {
        return this.go_user;
    }
    set_user(u) {
        this.go_user = u;
    }
}
exports.store = new go_data();
const element_exists = (n) => {
    const ele = document.getElementById(n);
    return ele == null ? false : true;
};
exports.element_exists = element_exists;
