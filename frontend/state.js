export const MAX_UPLOAD_BYTES = 10 * 1024 * 1024;
export const ACCEPTED_TYPES = ["image/jpeg", "image/png"]

export const state = {
    selectedFile: null,
    previewUrl: null,
    validationError: null,
    isSubmitting: false, //Guards against overlapping page submissions
    uploadError: null,
    storedImage: null, //Set once POST/v1/images succeeds
    job: null,
};

export function setSelectedFile(file, previewUrl){
    state.selectedFile = file;
    state.previewUrl = previewUrl;
    state.validationError = null;
    state.uploadError = null;
    state.storedImage = null;
    state.job = null;
}

export function clearSelection(){
    if (state.previewUrl) URL.revokeObjectURL(state.previewUrl);
    state.selectedFile = null;
    state.previewUrl = null;
    state.validationError = null;
    state.uploadError = null;
    state.storedImage = null;
    state.job = null;
}

export function setValidationError(message){
    state.validationError = message;
}

// Client-side check only
export function validateFileLocally(file){
    if(!ACCEPTED_TYPES.includes(file.type)) {
        return "Please choose a JPEG or PNG image.";
    }
    if(file.size > MAX_UPLOAD_BYTES) {
        return "That file is larger than 10 MB";
    }
    return null;
}
