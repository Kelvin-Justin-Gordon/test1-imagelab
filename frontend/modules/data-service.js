// Network calls live here so app.js never touches fetch() directly

const API_BASE = "/v1";

export class ApiError extends Error{
    constructor(message, status){
        super(message);
        this.status = status;
    }
}

export async function uploadImage(file){
    const formData = new FormData();
    formData.append("image", file);

    const response = await fetch(`${API_BASE}/images`, {
        method: "POST",
        body: formData,
    });

    const body = await response.json().catch(()=>null);

    if(!response.ok){
        const message = body?.error
        ? typeof body.error === "string"  ? body.error: Object.values(body.error).join(", "):`Upload failed (${response.status})`;
        throw new ApiError(message, response.status);
    }
    return body.image;
}