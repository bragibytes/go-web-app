const post_creator = "post-creator"
const vote_box = "vote-box"

const post_creator_element = document.getElementById(post_creator) as HTMLFormElement;
const vote_box_element = document.getElementById(vote_box) as HTMLDivElement;

import { 
    create_post,
 } from "../api";
import { post} from "../interfaces";
import { element_exists } from "./config";

const post_creator_handler = () => {

    const title = () => {
        return post_creator_element.querySelector('[name="title"]') as HTMLInputElement;
    }
    const content = () => {
        return post_creator_element.querySelector('[name="content"]') as HTMLInputElement;
    }
    const author = () => {
        return post_creator_element.querySelector('[name="author"]') as HTMLInputElement;
    }
    const on_submit = async (e:Event) => {
        e.preventDefault();
        console.log(title().value, content().value, author().value);
        const data:post = {
            title: title().value,
            content: content().value,
        }
        create_post(data, author().value)
    }
    post_creator_element.addEventListener('submit', on_submit)
}

const run = () => {
    if(element_exists(post_creator)){
        post_creator_handler()
    }
}

export default run;