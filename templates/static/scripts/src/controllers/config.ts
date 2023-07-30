
export const element_exists = (id:string):boolean => {
    const element = document.getElementById(id);
    return element !== null ? true:false;
}

export const chemical_x = ():any => {
    if (element_exists('json-data')) {
        const hidden_package = document.getElementById("json-data") as HTMLInputElement;
        const dataStrBase64 = hidden_package.value as string;
        const dataStr = atob(dataStrBase64);
        const dataObj = JSON.parse(dataStr);
        return dataObj;
    }
}
