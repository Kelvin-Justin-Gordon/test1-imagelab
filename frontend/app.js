import {
    state,
    setSelectedFile,
    clearSelection,
    setValidationError,
    validateFileLocally,
} from "./state.js";

import { render } from "./render.js";
import { uploadImage, ApiError } from "./modules/data-service";

const fileInput = document.getElementById("file-input");
const changeFileBtn = document.getElementById("change-file-btn");
const processBtn = document.getElementById("process-btn");

//Selecting a file only creates a local preview. No server job, until the user clicks "Process image".
fileInput.addEventListener("change", ()=> {
    const file = fileInput.files[0];
    if(!file) return;

    const previewUrl = URL.createdObjectURL(file);
    setSelectedFile(file, previewUrl);

    const error = validateFileLocally(file);
    if(error) setValidationError(error);

    render();
});

//Lets the user pick a different file before submitting.
changeFileBtn.addEventListener("click", ()=>{
    clearSelection();
    fileInput.value = "";
    render();
});

processBtn.addEventListener("click", SubmitImage);

async function submitImage(){
    if(state.isSubmitting) return;
    if (!state.selectedFile || state.validationError) return;

    state.isSubmitting = true;
    state.uploadError = null;
    render();

    try {
        const image = await uploadImage(state.selectedFile);
        state.storedImage = image;
    } catch(err){
        state.uploadError = err instanceof ApiError ? err.message : "Upload failed. Please try again";
    } finally{
        state.isSubmitting = false;
        render();
    }
}

render();
