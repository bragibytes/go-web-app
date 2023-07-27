import { vote } from "../interfaces";
import { vote_on_post } from "../api";
import { element_exists } from "./config";

const vote_button = "vote"

const vote_button_elements = document.getElementsByClassName(vote_button)

const run = () => {
   [...vote_button_elements].forEach(button => {
        button.addEventListener("click", () => {
            const valStr = button.getAttribute("value") as string
            const val = JSON.parse(valStr) as number
            const parentID = button.getAttribute("parent") as string
            const data:vote = {
                value:val,
                _parent:parentID
            }
            vote_on_post(data)
        })
    })
   
}

export default run;