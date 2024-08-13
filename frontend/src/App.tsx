
import './App.css'
import { Route, Routes } from 'react-router-dom'
import { Home } from './pages/Home'
import { Navbar } from './components/Navbar'
import { How } from './pages/How'
import { Products } from './pages/Products'
import { StoreProvider, useStoreContext } from './utils/storeContext'
import { useEffect } from 'react'

function App() {
  const { addProducts, addPrices } = useStoreContext()

  useEffect(() => {
    const fetchData = async () => {
      try {
        // Fetch products
        const productsResponse = await fetch('/api/products');
        const productsData = await productsResponse.json();
        addProducts(productsData);
        console.log(productsData)
        // Fetch prices
        const pricesResponse = await fetch('/api/prices');
        const pricesData = await pricesResponse.json();
        addPrices(pricesData);
        console.log(pricesData)

      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, []);
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/how" element={<How />} />
        <Route path='/juices' element={<Products />} />
      </Routes>
    </>
  )
}

export default App
