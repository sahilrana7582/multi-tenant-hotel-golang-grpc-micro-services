import { BrowserRouter, useRoutes } from 'react-router-dom';
import appRoutes from './routes/routes';
import { MantineProvider } from '@mantine/core';
import  { Toaster } from 'react-hot-toast';
import { Provider } from 'react-redux';
import { store } from './store/store';

function AppRoutes() {
  const routes = useRoutes(appRoutes);
  return routes;
}

function App() {
  return (

    <Provider store={store}>
      <MantineProvider>
        <BrowserRouter>
          <Toaster />
          <AppRoutes />
        </BrowserRouter>
      </MantineProvider>
    </Provider>
  );
}

export default App
