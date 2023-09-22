import Image from 'next/image';
import styles from './page.module.css';

import turtleImg from 'public/turtle01.jpg';

export default function Home() {
  return (
    <>
      <h1>水平思考の間へようこそ！</h1>
      <Image src={turtleImg} alt='turtle' />
    </>
  );
}
