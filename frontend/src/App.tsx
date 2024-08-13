
import './App.css'
import { Route, Routes } from 'react-router-dom'
import { Home } from './pages/Home'
import { Navbar } from './components/Navbar'
import { How } from './pages/How'
import { Products } from './pages/Products'
import { StoreProvider } from './utils/storeContext'

function App() {
  return (
    <>
      <StoreProvider>
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/how" element={<How />} />
          <Route path='/juices' element={<Products />} />
        </Routes>
      </StoreProvider>
    </>
  )
}

export default App
