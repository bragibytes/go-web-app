
export const element_exists = (id:string):boolean => {
    const element = document.getElementById(id);
    return element !== null ? true:false;
}