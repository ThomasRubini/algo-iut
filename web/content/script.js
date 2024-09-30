function initWasm() {
    if (!WebAssembly.instantiateStreaming) {
        // polyfill
        WebAssembly.instantiateStreaming = async (resp, importObject) => {
          const source = await (await resp).arrayBuffer();
          return await WebAssembly.instantiate(source, importObject);
        };
    }

    const go = new Go();

    let mod, inst;

    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
    async result => {
        mod = result.module;
        inst = result.instance;
        await go.run(inst);
    }
    );
}

var leftEditorInst = null;
var rightEditorInst = null;

async function initMonaco() {
    require.config({ paths: { 'vs': 'https://unpkg.com/monaco-editor@latest/min/vs' }});
    window.MonacoEnvironment = { getWorkerUrl: () => proxy };
    
    let proxy = URL.createObjectURL(new Blob([`
        self.MonacoEnvironment = {
            baseUrl: 'https://unpkg.com/monaco-editor@latest/min/'
        };
        importScripts('https://unpkg.com/monaco-editor@latest/min/vs/base/worker/workerMain.js');
    `], { type: 'text/javascript' }));
    
    let promise = new Promise((resolve, reject) => {
        require(["vs/editor/editor.main"], function () {
            leftEditorInst = monaco.editor.create(document.getElementById('leftEditorTag'), {
                language: 'javascript',
                theme: 'vs-dark'
            });
            rightEditorInst = monaco.editor.create(document.getElementById('rightEditorTag'), {
                language: 'javascript',
                theme: 'vs-dark'
            });    
            resolve();
        });
    });
    await promise;
}