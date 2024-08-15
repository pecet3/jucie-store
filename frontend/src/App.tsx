
import './App.css'
import { Route, Routes } from 'react-router-dom'
import { Home } from './pages/Home'
import { Navbar } from './components/Navbar'
import { How } from './pages/How'
import { Products } from './pages/Products'
import { useStoreContext } from './utils/storeContext'
import { useEffect, useState } from 'react'
import { Product } from './pages/Product'

function App() {
  const { addProducts, addPrices } = useStoreContext()
  const [isLoading, setIsLoading] = useState(true)
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
        setIsLoading(false)
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, []);
  return (
    <>
      <Navbar />
      {isLoading
        ? <p>Loading</p>
        :
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/how" element={<How />} />
          <Route path='/juices' element={<Products />} />
          <Route path="/juices/:id" element={<Product />} />
        </Routes>
      }
    </>
  )
}

export default App
