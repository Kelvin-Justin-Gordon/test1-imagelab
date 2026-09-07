import { state } from "./state.js";

const els = {
    dropLabel: document.getElementById("file-drop-label"),
    previewArea: document.getElementById("preview-area"),
    previewImage: document.getElementById("preview-image"),
    previewFilename: document.getElementById("preview-filename"),
    previewSize: document.getElementById("preview-size"),
    previewType: document.getElementById("preview-type"),
    validationError: document.getElementById("validation-error"),
    processBtn: document.getElementById("process-btn"),
    jobPanel: document.getElementById("job-panel"),
    jobIdLine: document.getElementById("job-id-line"),
    jobStatusLine: document.getElementById("job-status-line"),
};

function formatBytes(bytes){
    if(bytes < 1024) return `${bytes} B`;
    if(bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(1)} KB`;
    return `${(bytes / (1024 * 1024)).toFixed(2)} MB`;
}

//Re-renders the whole panel from state
export function render(){
    renderSelection();
    renderValidationError();
    renderProcessButton();
    renderJobPanel();
}

function renderSelection(){
    const hasFile = Boolean(state.selectedFile);
    els.previewArea.hidden = !hasFile;
    els.dropLabel.textContent = hasFile ? "Selection made - choose another file to replace it" : "Choose a JPEG or PNG (max 10 MB)";

    if (!hasFile) return;

    els.previewImage.src = state.previewUrl;
    els.previewFilename.textContent = state.selectedFile.name;
    els.previewSize.textContent = formatBytes(state.selectedFile.size);
    els.previewType.textContent = state.selectedFile.type || "unknown";
}

function renderValidationError(){
    const message = state.validationError || state.uploadError;
    els.validationError.hidden = !message;
    els.validationError.textContent = message || "";
}

function renderProcessButton(){ //disabled until an acceptable file is selected
    const canSubmit = Boolean(state.selectedFile) && !state.validationError && !state.isSubmitting;
    els.processBtn.disabled = !canSubmit;
    els.processBtn.textContent = state.isSubmitting ? "Uploading..." : "Process image";
}

function renderJobPanel(){
    if (!state.storedImage){
        els.jobPanel.hidden = true;
        return;
    }
    els.jobPanel.hidden = false;
    els.jobIdLine.textContent = `Image stored: ${state.storedImage.id}`;
}