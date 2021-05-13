const testRequest = () => {
    fetch('https://example.com')
        .then(res => console.log(res))
        .catch(err => console.error(err));
};

export {testRequest};
