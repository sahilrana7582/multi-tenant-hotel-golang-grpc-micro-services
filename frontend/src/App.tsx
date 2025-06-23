import { BrowserRouter, useRoutes } from 'react-router-dom';
import appRoutes from './routes/routes';

function AppRoutes() {
  const routes = useRoutes(appRoutes);
  return routes;
}

function App() {
  return (
    <BrowserRouter>
      <AppRoutes />
    </BrowserRouter>
  );
}

export default App
