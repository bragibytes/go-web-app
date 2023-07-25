import { vote } from "../interfaces";
import { send_vote } from "../api";
import { element_exists } from "./config";

const vote_box = "vote-box"

const vote_box_elements = document.getElementsByClassName("vote-box")

const vote_box_handler = () => {
    const arr = Array.from(vote_box_elements)
    arr.forEach(box => {
        const buttons = box.querySelectorAll(".button");
        buttons.forEach(button => {
            button.addEventListener("click", (e) => {
                const is_upvote_string = button.getAttribute("is-upvote") as string;
                const is_upvote = JSON.parse(is_upvote_string) as boolean;
                const parent = button.getAttribute("parent") as string;
                const data:vote = {
                    is_upvote: is_upvote,
                }
                send_vote(data, parent)
            })
        })
    })
}

const run = () => {
    if(element_exists(vote_box)){
        vote_box_handler();
    }
}

export default run;