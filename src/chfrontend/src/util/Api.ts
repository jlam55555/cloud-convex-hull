import {API_URL} from "../env";

// declare types
interface PresignRequest {
    type: 'GET' | 'POST';
    key?: string;
}
interface PresignResponse {
    url: string;
    key?: string;
}

// get presigned request
const presignRequest = (request: PresignRequest): Promise<Response> => {
    return fetch(API_URL + '/presign', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(request)
    });
};

// put a
const putRequest = (file: any, url: any): Promise<Response> => {
    return fetch(url, {
        method: 'PUT',
        body: file
    });
};

export {presignRequest, putRequest};
