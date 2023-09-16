import Footer from './footer';
import Header from './header';
import Container from './container';

export default function Layout({ children }) {
  return (
    <>
      <Header />
      <main>
        <Container>{children}</Container>
      </main>
      <Footer />
    </>
  );
}
