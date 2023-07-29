import Swal from "sweetalert2"
import { comment, post } from "../interfaces";
import { create_comment, update_comment, delete_comment } from "../api";
import { element_exists, json_data } from "./config";

const creator = "comment-creator";
const config = "comment-config";

const config_elements = document.getElementsByClassName(config)


const comment_creator_handler = () => {
    const data = json_data() as post
    const on_click = (e:Event) => {
        e.preventDefault()
        Swal.fire({
            title: "Create Comment",
            html: `
                <div class="container">
                    <div class="row">
                        <textarea id="content" class="swal2-input" placeholder="Content..."></textarea>
                    </div>
                </div>
            `,
            confirmButtonText: "Comment",
            showCancelButton: true,
            preConfirm: async () => {
                const content_input = Swal.getPopup()!.querySelector("#content") as HTMLInputElement;
                // Retrieve user input and handle data
                const content: string = content_input ? content_input.value:""
                // Do something with the newUsername and newEmail, e.g., send it to the server
                const new_comment:comment = {
                    content:content,
                    _parent: data._id
                }
                console.log(new_comment)
                console.log(data._id)
                create_comment(new_comment)
                
            },
        });
    }

    document.getElementById(creator)!.addEventListener("click", on_click)
}



const run = () => {
    if(element_exists(creator)) {
        comment_creator_handler()
    }
    [...config_elements].forEach((ele)=>{
        const comment_id = ele.getAttribute('comment_id') as string
        const current_content = ele.getAttribute('current_content') as string

        ele.addEventListener('click', (e:Event) => {
            Swal.fire({
                title: 'Update or delete your comment. Click anywhere outside this box to cancel.',
                html:`
                <div class="container">
                    <div class="row">
                        <textarea id="new-comment-content" class="swal2-input">${current_content}</textarea>
                    </div>
                </div>
                `,
                showCancelButton: true,
                confirmButtonColor: '#3085d6',
                cancelButtonColor: '#d33',
                confirmButtonText: 'Update',
                cancelButtonText: 'Delete',
                reverseButtons: true,
                preConfirm: () => {
                    const content_input = Swal.getPopup()?.querySelector('#new-comment-content') as HTMLTextAreaElement
                    const new_content = content_input.value
                    console.log("going to update comment with new content of --- "+new_content)
                    update_comment(comment_id, new_content.trim())
                }
              }).then((result) => {
                if (result.dismiss === Swal.DismissReason.cancel) {
                  // Delete comment
                  console.log("deleting comment ", comment_id)
                  delete_comment(comment_id)
                }
              })
        })
    })
}

export default run