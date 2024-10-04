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

function languageSyntax() {
// Create your own language definition here
// You can safely look at other samples without losing modifications.
// Modifications are not saved on browser refresh/close though -- copy often!
return {
    keywords: [
      // base
      'algorithme', 'debut', 'fin',
      // loops
      'faire', 'ffaire', 'boucle', 'fboucle', 'pour', 'jusqua', 'repeter', 'tant_que',
      // control flow
      'si', 'sinon', 'fsi',
      // other
      'declarer', 'ligne_suivante',
    ],
  
    typeKeywords: [
      'boolean', 'double', 'byte', 'int', 'short', 'char', 'void', 'long', 'float'
    ],
  
    operators: [
      '<-',
      'vaut', 'ne_veut_pas', '<', '>', '<=', '>=',
      'variant_de'
    ],
  
    // we include these common regular expressions
    symbols:  /[=><!~?:&|+\-*\/\^%]+/,
  
    // C# style strings
    escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  
    // The main tokenizer for our languages
    tokenizer: {
      root: [
        // identifiers and keywords
        [/[a-z_$][\w$]*/, { cases: { '@typeKeywords': 'keyword',
                                     '@keywords': 'keyword',
                                     '@default': 'identifier' } }],
        [/[A-Z][\w\$]*/, 'type.identifier' ],  // to show class names nicely
  
        // whitespace
        { include: '@whitespace' },
  
        // delimiters and operators
        [/[{}()\[\]]/, '@brackets'],
        [/[<>](?!@symbols)/, '@brackets'],
        [/@symbols/, { cases: { '@operators': 'operator',
                                '@default'  : '' } } ],
  
        // @ annotations.
        // As an example, we emit a debugging log message on these tokens.
        // Note: message are supressed during the first load -- change some lines to see them.
        [/@\s*[a-zA-Z_\$][\w\$]*/, { token: 'annotation', log: 'annotation token: $0' }],
  
        // numbers
        [/\d*\.\d+([eE][\-+]?\d+)?/, 'number.float'],
        [/0[xX][0-9a-fA-F]+/, 'number.hex'],
        [/\d+/, 'number'],
  
        // delimiter: after number because of .\d floats
        [/[;,.]/, 'delimiter'],
  
        // strings
        [/"([^"\\]|\\.)*$/, 'string.invalid' ],  // non-teminated string
        [/"/,  { token: 'string.quote', bracket: '@open', next: '@string' } ],
  
        // characters
        [/'[^\\']'/, 'string'],
        [/(')(@escapes)(')/, ['string','string.escape','string']],
        [/'/, 'string.invalid']
      ],
  
      string: [
        [/[^\\"]+/,  'string'],
        [/@escapes/, 'string.escape'],
        [/\\./,      'string.escape.invalid'],
        [/"/,        { token: 'string.quote', bracket: '@close', next: '@pop' } ]
      ],
  
      whitespace: [
        [/\/\/.*$/,    'comment'],
      ],
    },
  };
  
  
}

function initLanguage() {
    monaco.languages.register({ id: 'algo-iut' });
    let syntax = languageSyntax();
    monaco.languages.setMonarchTokensProvider('algo-iut', syntax);
    monaco.languages.registerCompletionItemProvider('algo-iut', {
        provideCompletionItems: function (model, position) {
            const suggestions = [
                ...syntax.keywords.map(k => {
                    return {
                        label: k,
                        kind: monaco.languages.CompletionItemKind.Keyword,
                        insertText: k
                    }
                })
            ]
            return { suggestions: suggestions };
        }
    });
}

async function initMonaco() {
    require.config({ paths: { 'vs': 'https://unpkg.com/monaco-editor@latest/min/vs' } });
    window.MonacoEnvironment = { getWorkerUrl: () => proxy };

    let proxy = URL.createObjectURL(new Blob([`
        self.MonacoEnvironment = {
            baseUrl: 'https://unpkg.com/monaco-editor@latest/min/'
        };
        importScripts('https://unpkg.com/monaco-editor@latest/min/vs/base/worker/workerMain.js');
    `], { type: 'text/javascript' }));

    let promise = new Promise((resolve, reject) => {
        require(["vs/editor/editor.main"], function () {
            initLanguage();

            leftEditorInst = monaco.editor.create(document.getElementById('leftEditorTag'), {
                language: 'algo-iut',
                theme: 'vs-dark'
            });
            rightEditorInst = monaco.editor.create(document.getElementById('rightEditorTag'), {
                language: 'cpp',
                theme: 'vs-dark',
                readOnly: true
            });
            resolve();
        });
    });
    await promise;
}