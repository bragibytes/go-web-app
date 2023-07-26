
export const element_exists = (id:string):boolean => {
    const element = document.getElementById(id);
    return element !== null ? true:false;
}

export const json_data = ():any => {
    const ele = document.getElementById("json-data") as HTMLInputElement
    const strData = ele.value as string
    const obj = JSON.parse(strData)
    return obj
}