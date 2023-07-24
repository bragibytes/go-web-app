"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const authentication_1 = __importDefault(require("./controllers/authentication"));
const config_1 = require("./controllers/config");
if ((0, config_1.element_exists)("hidden-package")) {
    const x = document.getElementById("hidden-package");
    const s = x.value;
    const u = JSON.parse(s);
    config_1.store.set_user(u);
}
(0, authentication_1.default)();
// users()
