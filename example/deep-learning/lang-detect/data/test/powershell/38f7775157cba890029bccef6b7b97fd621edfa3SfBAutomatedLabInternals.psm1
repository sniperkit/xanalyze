function Read-Choice
{ 
    param(
        [Parameter(Mandatory)]
        [String[]]$ChoiceList, 

        [Parameter(Mandatory)]
        [String]$Caption,
        
        [String]$Message,

        [int]$Default = 0
    )
    
    if (-not $Message) { $Message = $Caption }

    $choices = New-Object System.Collections.ObjectModel.Collection[System.Management.Automation.Host.ChoiceDescription]

    $choiceList | ForEach-Object { $choices.Add((New-Object "System.Management.Automation.Host.ChoiceDescription" -ArgumentList $_)) }

    $Host.UI.PromptForChoice($Caption, $Message, $choices, $Default) 
}

function Read-HashTable
{ 
    param(
        [Parameter(Mandatory)]
        [String[]]$ChoiceList, 

        [Parameter(Mandatory)]
        [String]$Caption,
        
        [String]$Message,

        [int]$Default = 0
    )
    
    #if (-not $Message) { $Message = $Caption }

    $fields = New-Object System.Collections.ObjectModel.Collection[System.Management.Automation.Host.FieldDescription]

    $choiceList | ForEach-Object { $fields.Add((New-Object System.Management.Automation.Host.FieldDescription -ArgumentList $_)) }

    $Host.UI.Prompt($Caption, $Message, $fields)    
}