class Program
{
    static void Main()
    {
        ILogger logger = new FileLogger("mylog.txt");
        // ILogger logger = new ConsoleLogger();
        // igual a acima  - ConsoleLogger logger = new ConsoleLogger();
        BankAccount account1 = new BankAccount("Raquel", 100, logger);
        BankAccount account2 = new BankAccount("Mariana", 100, logger);

        account1.Deposit(-100);
        account2.Deposit(150);

        Console.WriteLine(account2.Balance);

        
    }
}

class FileLogger : ILogger
{
    private readonly string filePath;

    //ctor - construtor publico
    public FileLogger(string filePath)
    {
        this.filePath = filePath;
        this.filePath = filePath;
    }
    public void Log(string message)
    {
        //append - serve para adicionar ao final do conteudo do arquivo 
        File.AppendAllText("filePath.txt", $"{message}{Environment.NewLine}"); //.newlinw = \n
    }
}
class ConsoleLogger : ILogger //utiizar ctrl + . para implementar a interface
{
}

interface ILogger //convenção, sempre iniciar o nome com "I"
{
    //membros de interface não declara acessibilidade, todos são publicos
    void Log(string message)
    {
        Console.WriteLine($"LOGGER: {message}");
    }
}


class BankAccount
{
    private string name;
    private readonly ILogger logger; //readonly - só vai poder atribuir valor dentro de um construtor
    private decimal balance;//saldo da conta

    public decimal Balance
    {
        get; private set;
    }


    public BankAccount(string name, decimal balance, ILogger logger)// ctrl+. no parametro = create and assign fiel(coloca this. automatico)
    {
        if (string.IsNullOrWhiteSpace(name))
        {
            throw new ArgumentException("Nome inválido", nameof(name));//nameof transforma em string
        }
        if (balance < 0)
        {
            throw new Exception("Saldo não pode ser negativado.");
        }
        this.name = name;
        Balance = balance;
        this.logger = logger;
    }

    public void Deposit(decimal amount)
    {
        if (amount <= 0)
        {
            logger.Log($"Não é possível depositar {amount} na conta de {name}."); //log - imprimir no console
            // throw new Exception("Saldo não pode ser negativado.");
            return; //finaliza a execução do método
        }
        Balance += amount;
    }
}
