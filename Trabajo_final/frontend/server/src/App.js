import './App.css';
import { Cabecera } from './components/Cabecera';
import { SegundoComponente } from './components/SegundoComponente';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <a name="tope"></a> 
        <Cabecera />
      </header>
      <body className="App-body"> 
      <div className="Principal">
       <SegundoComponente /> 
      </div>
        </body>

    </div>
    
  );
}

export default App;
