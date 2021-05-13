import {API_URL} from "../env";

// declare types
interface PresignRequest {
    type: 'GET' | 'PUT';
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

// put a file onto s3
const putRequest = (file: any, url: any): Promise<Response> => {
    return fetch(url, {
        method: 'PUT',
        body: file
    });
};

const convexHullRequest = (key: any): Promise<Response> => {
    return fetch(API_URL + '/convexhull', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({key})
    });
};

export {presignRequest, putRequest, convexHullRequest};
