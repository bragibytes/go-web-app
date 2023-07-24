"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.notify_modal = exports.notify = void 0;
const notie_1 = __importDefault(require("notie"));
const sweetalert2_1 = __importDefault(require("sweetalert2"));
const notify = (msg, msgType) => {
    notie_1.default.alert({
        type: msgType,
        text: msg,
    });
};
exports.notify = notify;
const notify_modal = (icon, title, text, footer) => {
    return sweetalert2_1.default.fire({
        icon: icon,
        title: title,
        text: text,
        footer: footer
    });
};
exports.notify_modal = notify_modal;
