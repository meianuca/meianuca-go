function calculaPraia() {
  var ondulacao = comboOndulacao.options[comboOndulacao.selectedIndex].value;
  var vento = comboVento.options[comboVento.selectedIndex].value;
  var praiaResultado = 'Sei não';

  if(ondulacao == 'nordeste'){
    if(vento == 'sul' || vento == 'sudeste'){
      praiaResultado = 'Barra da Lagoa';
    } else if (vento == 'sudoeste'){
      praiaResultado = 'Santinho, Caldeirão, Barra da Lagoa';
    } else if (vento == 'norte' || vento == 'nordeste'){
      praiaResultado = 'Lagoinha do Leste, Galheta Norte, Galheta, Joaquina, Moçambique, Mole, Santinho';
    }
  } else if (ondulacao == 'sudeste'){
    if(vento == 'sul' || vento == 'sudoeste'){
      praiaResultado = 'Riozinho, Barra da Lagoa';
    } else if(vento == 'norte' || vento == 'nordeste'){
      praiaResultado = 'Morro das Pedras, Novo Campeche, Lomba do Sabão, Galheta, Galheta Norte, Joaquina, Moçambique, Mole, Pico da Cruz, Santinho, Igrejinha, Lagoinha do Leste';
    } else if (vento == 'sudeste'){
      praiaResultado = 'Barra da Lagoa';
    } else if (vento == 'noroeste'){
      praiaResultado = 'Morro das Pedras';
    } else if (vento == 'oeste'){
      praiaResultado = 'Riozinho';
    }
  } else if (ondulacao == 'sul'){
    if(vento == 'norte' || vento == 'nordeste'){
      praiaResultado = 'Açores, Joaquina, Moçambique, Santinho, Mole, Solidão, Naufragados, Galheta Norte, Lagoinha do Leste';
    } else if(vento == 'sudoeste' || vento == 'oeste'){
      praiaResultado = 'Campeche';
    } else if(vento == 'nordeste'){
      praiaResultado = 'Morro das Pedras';
    } 
  } else if (ondulacao == 'leste'){
    if(vento == 'sudoeste'){
      praiaResultado = 'Armação, Brava, Caldeirão, Riozinho, Ingleses, Matadeiro';
    } else if(vento == 'noroeste'){
      praiaResultado = 'Morro das Pedras';
    } else if(vento == 'sul'){
      praiaResultado = 'Brava, Ingleses, Matadeiro';
    } else if(vento == 'norte' || vento == 'nordeste'){
      praiaResultado = 'Pico da Cruz, Igrejinha, Novo Campeche, Lomba do Sabão, Santinho';
    } else if(vento == 'oeste'){
      praiaResultado = 'Riozinho';
    }
  } else if (ondulacao == 'nordeste'){
    if(vento == 'sul' || vento == 'sudeste'){
      praiaResultado = 'Barra da Lagoa';
    } else if(vento == 'sudoeste'){
      praiaResultado = 'Santinho, Caldeirão, Barra da Lagoa';
    } else if(vento == 'norte' || vento == 'nordeste'){
      praiaResultado = 'Lagoinha do Leste, Galheta Norte, Galheta, Joaquina, Moçambique, Mole, Santinho';
    }
  }

  var divPraia = document.getElementById("divPraia");
  divPraia.innerText = praiaResultado;
}