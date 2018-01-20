/*
 * File automatically generated by
 * gengen 1.0 by Lorenzo Bettini 
 * http://www.gnu.org/software/gengen
 */

#include "multiple_option.h"

void
multiple_option_gen_class::generate_multiple_option(ostream &stream, unsigned int indent)
{
  string indent_str (indent, ' ');
  indent = 0;

  if (has_short_option)
    {
      stream << "case '";
      generate_string (short_option, stream, indent + indent_str.length ());
      stream << "':	/* ";
      generate_string (option_comment, stream, indent + indent_str.length ());
      stream << ".  */";
      stream << "\n";
      stream << indent_str;
    }
  else
    {
      stream << "/* ";
      generate_string (option_comment, stream, indent + indent_str.length ());
      stream << ".  */";
      stream << "\n";
      stream << indent_str;
      generate_string (gen_else, stream, indent + indent_str.length ());
      stream << "if (strcmp (long_options[option_index].name, \"";
      generate_string (long_option, stream, indent + indent_str.length ());
      stream << "\") == 0)";
      stream << "\n";
      stream << indent_str;
      stream << "{";
      stream << "\n";
      stream << indent_str;
    }
  stream << "  local_args_info.";
  generate_string (option_var_name, stream, indent + indent_str.length ());
  stream << "_given++;";
  stream << "\n";
  stream << indent_str;
  stream << "\n";
  stream << indent_str;
  if (option_has_type)
    {
      stream << "  multi_token = get_multiple_arg_token(optarg);";
      stream << "\n";
      stream << indent_str;
      stream << "  multi_next = get_multiple_arg_token_next (optarg);";
      stream << "\n";
      stream << indent_str;
      stream << "\n";
      stream << indent_str;
      stream << "  while (1)";
      stream << "\n";
      stream << indent_str;
      stream << "    {";
      stream << "\n";
      stream << indent_str;
      if (option_has_values)
        {
          stream << "      if ((found = check_possible_values((multi_token ? multi_token : ";
          if (( defaultval != "" ))
            {
              stream << "\"";
              generate_string (defaultval, stream, indent + indent_str.length ());
              stream << "\" ";
            }
          else
            {
              stream << "0 ";
            }
          stream << "), ";
          generate_string (option_values, stream, indent + indent_str.length ());
          stream << ")) < 0)";
          stream << "\n";
          stream << indent_str;
          stream << "        {";
          stream << "\n";
          stream << indent_str;
          stream << "          fprintf (stderr, \"%s: %s argument, \\\"%s\\\", for option `--";
          generate_string (long_option, stream, indent + indent_str.length ());
          stream << "'";
          if (has_short_option)
            {
              stream << " (`-";
              generate_string (short_option, stream, indent + indent_str.length ());
              stream << "')";
            }
          stream << "%s\\n\", ";
          generate_string (package_var_name, stream, indent + indent_str.length ());
          stream << ", (found == -2) ? \"ambiguous\" : \"invalid\", multi_token, (additional_error ? additional_error : \"\"));";
          stream << "\n";
          stream << indent_str;
          stream << "          free (multi_token);";
          stream << "\n";
          stream << indent_str;
          stream << "          goto failure;";
          stream << "\n";
          stream << indent_str;
          stream << "        }";
          stream << "\n";
          stream << indent_str;
          stream << "\n";
          stream << indent_str;
        }
      stream << "      ";
      generate_string (option_var_name, stream, indent + indent_str.length ());
      stream << "_new = (struct ";
      generate_string (type, stream, indent + indent_str.length ());
      stream << "_list *) malloc (sizeof (struct ";
      generate_string (type, stream, indent + indent_str.length ());
      stream << "_list));";
      stream << "\n";
      stream << indent_str;
      indent = 6;
      stream << "      ";
      generate_string (option_var_name, stream, indent + indent_str.length ());
      stream << "_new->next = ";
      generate_string (option_var_name, stream, indent + indent_str.length ());
      stream << "_list;";
      indent = 0;
      stream << "\n";
      stream << indent_str;
      indent = 6;
      stream << "      ";
      generate_string (option_var_name, stream, indent + indent_str.length ());
      stream << "_list = ";
      generate_string (option_var_name, stream, indent + indent_str.length ());
      stream << "_new;";
      indent = 0;
      stream << "\n";
      stream << indent_str;
      indent = 6;
      stream << "      ";
      generate_string (update_arg, stream, indent + indent_str.length ());
      indent = 0;
      stream << "\n";
      stream << indent_str;
      stream << "\n";
      stream << indent_str;
      stream << "      if (multi_next)";
      stream << "\n";
      stream << indent_str;
      stream << "        {";
      stream << "\n";
      stream << indent_str;
      stream << "          multi_token = get_multiple_arg_token(multi_next);";
      stream << "\n";
      stream << indent_str;
      stream << "          multi_next = get_multiple_arg_token_next (multi_next);";
      stream << "\n";
      stream << indent_str;
      stream << "          local_args_info.";
      generate_string (option_var_name, stream, indent + indent_str.length ());
      stream << "_given++;";
      stream << "\n";
      stream << indent_str;
      stream << "        }";
      stream << "\n";
      stream << indent_str;
      stream << "      else";
      stream << "\n";
      stream << indent_str;
      stream << "        break;";
      stream << "\n";
      stream << indent_str;
      stream << "    }";
      stream << "\n";
      stream << indent_str;
    }
  if (option_has_group)
    {
      stream << "  if (!args_info->";
      generate_string (option_var_name, stream, indent + indent_str.length ());
      stream << "_group)";
      stream << "\n";
      stream << indent_str;
      stream << "    {";
      stream << "\n";
      stream << indent_str;
      stream << "      args_info->";
      generate_string (option_var_name, stream, indent + indent_str.length ());
      stream << "_group = 1;";
      stream << "\n";
      stream << indent_str;
      stream << "      args_info->";
      generate_string (group_var_name, stream, indent + indent_str.length ());
      stream << "_group_counter += 1;";
      stream << "\n";
      stream << indent_str;
      stream << "    }";
      stream << "\n";
      stream << indent_str;
    }
  stream << "  break;";
  if (has_short_option)
    {
      

    }
  else
    {
      stream << "\n";
      stream << indent_str;
      stream << "}";
      stream << "\n";
      stream << indent_str;
    }
}
