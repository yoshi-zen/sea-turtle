import Image from 'next/image';
import styles from './page.module.css';

import turtleImg from 'public/turtle01.jpg';

export default function Home() {
  return (
    <>
      <div className={styles.hero_container}>
        <div className={styles.heroText_wrapper}>
          <h1>水平思考の間へようこそ！</h1>
        </div>
        <div className={styles.heroImg_wrapper}>
          <Image src={turtleImg} alt='turtle' width={500} priority />
        </div>
      </div>
    </>
  );
}
