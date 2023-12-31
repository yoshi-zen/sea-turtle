'use client';

import styles from '@/app/questions/page.module.css';
import Layout12 from '@/components/layout-12columns';
import WrapperShadowBox from '@/components/wrapper-shadow-box';

import { useEffect, useState } from 'react';

const qTitle = [
  { title: 'ウミガメのスープ', content: '「ウミガメのスープが飲みたい。」\n男は各地を旅行した。\nとある海辺のレストランで ウミガメのスープを出す店を見つけた。\n彼はウミガメのスープを注文し その日に自殺した。\n念願だったウミガメのスープが見つかったのに 彼はなぜ自殺したのだろうか？' },
  { title: 'クレイジータクシー', content: 'こんちゃ' },
  { title: 'ひどいスッピン', content: 'こんばんは' },
  { title: 'ゴンとチロ', content: 'おはよう' },
  { title: '過激なダイエット', content: 'おやすみ' },
  { title: '爆弾が仕掛けられた車', content: 'hogehoge' },
  { title: 'ある偉人の版画' },
  { title: '値段より大切なもの' },
  { title: '包帯を巻いたプログラマー' },

  // 'ウミガメのスープ',
  // 'クレイジータクシー',
  // 'ひどいスッピン',
  // 'ゴンとチロ',
  // '過激なダイエット',
  // '爆弾が仕掛けられた車',
  // 'ある偉人の版画',
  // '値段より大切なもの',
  // '包帯を巻いたプログラマー',
  // '刈り取り',
  // '知る人ぞ知るラーメン',
  // '男の覚悟',
  // '道を教えてください',
  // '振られましたね？',
  // '真逆の主人公',
  // 'プリクラ開発秘話',
  // 'カメラマンの願い',
  // '臓器提供',
  // '素敵なステーキ',
  // 'うんこを見ろ',
  // '親を食った男',
  // '箱の中',
  // 'テレビ壊しちゃった',
  // 'パーカーを着なさい',
];

export default function Questions() {
  const [selected, setSelected] = useState(null);
  const [content, setContent] = useState('');

  useEffect(() => {
    if (selected !== null) {
      setContent(qTitle[selected].content);
    }
  }, [selected]);

  const handleClick = (index) => {
    setSelected(index);
  };

  return (
    <Layout12>
      <div className={styles.left_container}>
        <h1 className={styles.pageTitle}>問題一覧</h1>
        {qTitle.map((data, index) => (
          <div key={index} className={selected === index ? styles.title_selected_inner : styles.title_inner} onClick={() => handleClick(index)}>
            <div className={styles.flex_inner}>
              <p style={{ width: '1.5em', textAlign: 'end' }}>{index + 1}.</p>
              <p>{data.title}</p>
            </div>
          </div>
        ))}
      </div>
      <div className={styles.center_container}>
        <div className={styles.flex_wrapper}>
          <div className={styles.question_inner}>
            {/* <p>「ウミガメのスープが飲みたい。」</p>
            <p>男は各地を旅行した。</p>
            <p>とある海辺のレストランで ウミガメのスープを出す店を見つけた。</p>
            <p>彼はウミガメのスープを注文し その日に自殺した。</p>
            <p>念願だったウミガメのスープが見つかったのに 彼はなぜ自殺したのだろうか？</p> */}
            {content.split('\n').map((data, index) => (
              <p key={index}>{data}</p>
            ))}
          </div>
          <div>
            <p>ここにチャット</p>
          </div>
        </div>
      </div>
      <div className={styles.right_container}>
        <WrapperShadowBox>
          <p>Contributer</p>
          <div className={styles.contributer_name_pic_inner}>
            <div style={{ width: '50px', height: '50px', borderRadius: '25px', backgroundColor: 'gray' }}></div>
            <div className={styles.contributer_name_button_inner}>
              <p style={{ fontWeight: '700', letterSpacing: 'normal' }}>Yoshi_zen</p>
              <button>Follow</button>
            </div>
          </div>
        </WrapperShadowBox>
        {/* <div className={styles.new_question_wrapper}>
          <h2>新着問題</h2>
          <p>ゴンとチロ</p>
        </div> */}
      </div>
    </Layout12>
  );
}
