import { vote } from "../interfaces";
import { send_vote } from "../api";
import { element_exists } from "./config";

const vote_button = "vote-button"

const vote_button_elements = document.getElementsByClassName("vote-button")

const run = () => {
   [...vote_button_elements].forEach(button => {
        button.addEventListener("click", () => {
            const upvote_string = button.getAttribute("upvote") as string;
            const is_upvote = JSON.parse(upvote_string) as boolean;
            const parent = button.getAttribute("parent") as string;
            const data:vote = {
                is_upvote: is_upvote,
            }
            send_vote(data, parent)
        })
    })
   
}

export default run;