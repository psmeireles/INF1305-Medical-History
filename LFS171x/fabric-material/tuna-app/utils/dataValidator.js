var axios = require("axios")

const validateCRM = (crm) => {
	const url = `https://www.consultacrm.com.br/api/index.php?tipo=crm&q=${crm}&chave=8194823720&destino=json`
	axios.get(url)
  .then(function (response) {
		// handle success
		data = response.data
		if(data.item.length == 0){
			console.log("Não existe nenhum médico cadastrado com esse CRM")
			return false
		}
		else if(data.item.length != 1){
			console.log("Mais de um médico identificado com esse CRM")
			return false
		}
		else{
			if(data.item[0].situacao == "Ativo"){
				console.log("CRM está ativo")
				return true
			}
			else{
				console.log("CRM inativo")
				return false
			}
		}
  })
  .catch(function (error) {
    // handle error
		console.log(error);
		return false
  })
}

const validateCNPJ = (cnpj) => {
	const url = `https://www.receitaws.com.br/v1/cnpj/${cnpj}`
	axios.get(url)
  .then(function (response) {
		// handle success
		data = response.data
		if(data.status == "ERROR"){
			console.log(data.message)
			return false
		}
		else{
			activities = data["atividade_principal"].concat(data["atividades_secundarias"])
			return activities.some(activity => {
				var code = parseInt(activity.code.substring(0, 2))
				if(code >= 86 && code <= 89){
					console.log("A empresa está relacionada à área de saúde")
					return true
				}
				else{
					console.log("A empresa não está relacionada à área de saúde")
					return false
				}
			});
		}
  })
  .catch(function (error) {
    // handle error
		console.log(error);
		return false
	})
}

export {validateCRM, validateCNPJ};